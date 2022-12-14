import { Body, Controller, Get, Param, Post, Query, Res } from '@nestjs/common';
import { GrpcMethod, RpcException } from '@nestjs/microservices';
import { ObjectId, Types } from 'mongoose';
import { EmbededService } from 'src/embeded/embeded.service';
import { ExamService } from 'src/exam/exam.service';
import { IS_MICROSERVICE } from 'src/main';
import {
  ThreadCreateBody,
  ThreadIdRequestBody,
  ThreadRequestAddAnswerBody,
  ThreadRequestBody,
  ThreadRequestCreateBody,
} from './thread.interface';
import { ThreadService } from './thread.service';

@Controller('thread')
export class ThreadController {
  constructor(
    private readonly ThreadService: ThreadService,
    private readonly ExamService: ExamService,
    private readonly EmbededService: EmbededService,
  ) {}

  @Post('/')
  @GrpcMethod('Thread')
  async createThread(
    @Res() response,
    @Body() reqBody: ThreadRequestCreateBody,
    grpcBody: ThreadRequestCreateBody,
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : reqBody;

    // check if exam exist, if not then create new exam
    const exams = await this.ExamService.findByCourseProperty({
      year: data.year,
      semester: data.semester,
      term: data.term,
    });

    let examId: ObjectId;

    if (exams.length) {
      // exam is existed
      examId = await this.ExamService.findIdByCourseProperty(data);
    } else {
      // exam isn't existed, then create exam before create thread
      const newExam = await this.ExamService.create({
        course_id: data.course_id,
        year: data.year,
        semester: data.semester,
        term: data.term,
      });

      examId = newExam._id;
    }

    const problemId = String(new Types.ObjectId());
    const problemVector = [0, 1, 2]; // will be generated from nlp

    // add problem's vector to embeded collection
    await this.EmbededService.create({
      problem_id: problemId,
      vector: problemVector,
    });

    // compare...
    const embededVectors = await this.EmbededService.find();
    const comparedVectors = embededVectors.map((vector) => {
      const similarity = 0.98; // compare logic here :: more than > 0.9
      if (similarity > 0.9) {
        return { _id: vector._id, similarity: similarity };
      }
    });

    if (comparedVectors.length > 0) {
      // add problem to existed thread
    } else {
      // create new thread
    }

    const createThreadBody: ThreadCreateBody = {
      exam_id: String(examId),
      course_id: data.course_id,

      upvoted: 0,
      downvoted: 0,
      problems: [
        {
          id: problemId,
          title: data.question_title,
          body: data.question,
          uploaded_user: data.uploaded_user,
          upvoted: 0,
          downvoted: 0,
        },
      ],
      answers: data.answer && [
        {
          id: String(new Types.ObjectId()),
          body: data.answer,
          upvoted: 0,
          downvoted: 0,
        },
      ],
    };

    const newThread = await this.ThreadService.createThread(createThreadBody);

    const threadId = newThread._id;

    // add thread_id to exam document
    await this.ExamService.addThreadId(threadId, examId);

    return IS_MICROSERVICE ? newThread : response.status(201).json(newThread);
  }

  @Get('/')
  @GrpcMethod('Thread')
  async getThread(
    @Res() response,
    @Query() reqBody: ThreadRequestBody,
    grpcBody: ThreadRequestBody,
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : reqBody;

    if (data.course_id && data.year && data.semester && data.term) {
      const examId = await this.ExamService.findIdByCourseProperty(data);

      if (examId) {
        const threads = await this.ThreadService.find({
          exam_id: String(examId),
          course_id: data.course_id,
        });

        return IS_MICROSERVICE
          ? { messages: threads }
          : response.status(200).json(threads);
      } else {
        return IS_MICROSERVICE
          ? { messages: [] }
          : response.status(404).json([]);
      }
    } else if (data.exam_id) {
      const thread = await this.ThreadService.findByExamId(data.exam_id);

      return IS_MICROSERVICE
        ? { messages: thread }
        : response.status(200).json(thread);
    } else {
      const threads = await this.ThreadService.findAll();
      return IS_MICROSERVICE
        ? { messages: threads }
        : response.status(200).json(threads);
    }
  }

  @Get('/:threadId')
  @GrpcMethod('Thread')
  async getThreadById(@Res() response, @Param('threadId') threadId) {
    const thread = await this.ThreadService.findOneById(threadId);

    if (thread) {
      return IS_MICROSERVICE ? thread : response.status(200).json(thread);
    } else {
      if (IS_MICROSERVICE) {
        throw new RpcException({ message: 'Not found', status: 5 });
      }
      return response.status(400).json([]);
    }
  }

  @Post('/upvote/:threadId')
  @GrpcMethod('Thread')
  async upvoteThread(
    @Res() response,
    @Param('threadId') threadId,
    grpcBody: ThreadIdRequestBody,
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : { thread_id: threadId };

    const thread = await this.ThreadService.findOneById(data.thread_id);

    const updatedUpvoted = thread.upvoted + 1;

    const updatedThread = await this.ThreadService.updateThreadById(threadId, {
      upvoted: updatedUpvoted,
    });

    return IS_MICROSERVICE
      ? updatedThread
      : response.status(200).json(updatedThread);
  }

  @Post('/downvote/:threadId')
  @GrpcMethod('Thread')
  async downvoteThread(
    @Res() response,
    @Param('threadId') threadId,
    grpcBody: ThreadIdRequestBody,
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : { thread_id: threadId };

    const thread = await this.ThreadService.findOneById(data.thread_id);

    const updatedDownvoted = thread.downvoted - 1;

    const updatedThread = await this.ThreadService.updateThreadById(threadId, {
      downvoted: updatedDownvoted,
    });

    return IS_MICROSERVICE
      ? updatedThread
      : response.status(200).json(updatedThread);
  }

  @Post('/problem/upvote/:problemId')
  @GrpcMethod('Thread')
  async upvoteProblem(
    @Res() response,
    @Param('problemId') problemId,
    grpcBody: { problem_id: string },
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : { problem_id: problemId };

    const problem = await this.ThreadService.findProblem(data.problem_id);

    const updatedUpvoted = problem.upvoted + 1;

    // update upvote
    await this.ThreadService.updateVoted(
      { problem: true, upvote: true },
      data.problem_id,
      updatedUpvoted,
    );

    const updatedProblem = await this.ThreadService.findProblem(
      data.problem_id,
    );

    return IS_MICROSERVICE
      ? updatedProblem
      : response.status(200).json(updatedProblem);
  }

  @Post('/problem/downvote/:problemId')
  @GrpcMethod('Thread')
  async downvoteProblem(
    @Res() response,
    @Param('problemId') problemId,
    grpcBody: { problem_id: string },
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : { problem_id: problemId };

    const problem = await this.ThreadService.findProblem(data.problem_id);

    const updatedDownvoted = problem.downvoted - 1;

    // update downvote
    await this.ThreadService.updateVoted(
      { problem: true, downvote: true },
      data.problem_id,
      updatedDownvoted,
    );

    const updatedProblem = await this.ThreadService.findProblem(
      data.problem_id,
    );

    return IS_MICROSERVICE
      ? updatedProblem
      : response.status(200).json(updatedProblem);
  }

  @Post('/answer/upvote/:answerId')
  @GrpcMethod('Thread')
  async upvoteAnswer(
    @Res() response,
    @Param('answerId') answerId,
    grpcBody: { answer_id: string },
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : { answer_id: answerId };

    const answer = await this.ThreadService.findAnswer(data.answer_id);

    const updatedUpvoted = answer.upvoted + 1;
    console.log(updatedUpvoted);

    // update downvote
    await this.ThreadService.updateVoted(
      { answer: true, upvote: true },
      data.answer_id,
      updatedUpvoted,
    );

    const updatedAnswer = await this.ThreadService.findAnswer(data.answer_id);

    return IS_MICROSERVICE
      ? updatedAnswer
      : response.status(200).json(updatedAnswer);
  }

  @Post('/answer/downvote/:answerId')
  @GrpcMethod('Thread')
  async downvoteAnswer(
    @Res() response,
    @Param('answerId') answerId,
    grpcBody: { answer_id: string },
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : { answer_id: answerId };

    const answer = await this.ThreadService.findAnswer(data.answer_id);

    const updatedDownvoted = answer.downvoted - 1;

    // update downvote
    await this.ThreadService.updateVoted(
      { answer: true, downvote: true },
      data.answer_id,
      updatedDownvoted,
    );

    const updatedAnswer = await this.ThreadService.findAnswer(data.answer_id);

    return IS_MICROSERVICE
      ? updatedAnswer
      : response.status(200).json(updatedAnswer);
  }

  @Post('/answer')
  @GrpcMethod('Thread')
  async addAnswer(
    @Res() response,
    @Body() reqBody: ThreadRequestAddAnswerBody,
    grpcBody: ThreadRequestAddAnswerBody,
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : reqBody;

    const thread = await this.ThreadService.findOneById(data.thread_id);

    const updatedAnswer = [
      ...thread.answers,
      {
        id: String(new Types.ObjectId()),
        body: data.body,
        upvoted: 0,
        downvoted: 0,
      },
    ];

    const updatedThread = await this.ThreadService.updateThreadById(
      data.thread_id,
      { answers: updatedAnswer },
    );

    return IS_MICROSERVICE
      ? updatedThread
      : response.status(201).json(updatedThread);
  }

  // @GrpcMethod()
  // async getThreadByTitle(data, metadata: any) {}

  // @GrpcMethod()
  // async addAnswerById(data: ThreadRequestAddAnswerBody, metadata: any) {
  //   const thread = await this.ThreadService.findOneById(data.thread_id);

  //   const threadId = thread._id;
  //   const answer = {
  //     body: data.body,
  //     upvoted: 0,
  //     downvoted: 0,
  //   };

  //   const updatedThread = await this.ThreadService.updateThreadById(threadId, {
  //     answers: [...thread.answers, answer],
  //   });

  //   return updatedThread;
  // }

  // @GrpcMethod()
  // async downvoteThread(data: ThreadIdRequestBody, metadata: any) {
  //   const thread = await this.ThreadService.findOneById(data.thread_id);

  //   const threadId = thread._id;
  //   const updatedDownvoted = thread.downvoted - 1;

  //   const updatedThread = await this.ThreadService.updateThreadById(threadId, {
  //     upvoted: updatedDownvoted,
  //   });

  //   return updatedThread;
  // }
}
