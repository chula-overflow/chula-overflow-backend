import { Body, Controller, Get, Param, Post, Query, Res } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { response } from 'express';
import { Types } from 'mongoose';
import { ExamPropertyRequestBody } from 'src/exam/exam.interface';
import { ExamService } from 'src/exam/exam.service';
import { IS_MICROSERVICE } from 'src/main';
import { threadId } from 'worker_threads';
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
  ) {}

  @Post('/')
  @GrpcMethod()
  async createThread(
    @Res() response,
    @Body() reqBody: ThreadRequestCreateBody,
    grpcBody: ThreadRequestCreateBody,
    metadata: any,
  ) {
    const data = IS_MICROSERVICE ? grpcBody : reqBody;

    const examId = await this.ExamService.findIdByCourseProperty(data);

    const createThreadBody: ThreadCreateBody = {
      exam_id: String(examId),
      course_id: data.course_id,

      upvoted: 0,
      downvoted: 0,
      problems: [
        {
          id: String(new Types.ObjectId()),
          title: 'generated from nlp',
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
  @GrpcMethod()
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

        return IS_MICROSERVICE ? threads : response.status(200).json(threads);
      } else {
        return IS_MICROSERVICE ? [] : response.status(404).json([]);
      }
    } else {
      return IS_MICROSERVICE ? [] : response.status(400).json([]);
    }
  }

  @Get('/:threadId')
  @GrpcMethod()
  async getThreadById(@Res() response, @Param('threadId') threadId) {
    const thread = await this.ThreadService.findOneById(threadId);

    if (thread) {
      return IS_MICROSERVICE ? thread : response.status(200).json(thread);
    } else {
      return IS_MICROSERVICE ? [] : response.status(400).json([]);
    }
  }

  @Post('/problem/upvote/:problemId')
  @GrpcMethod()
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
  @GrpcMethod()
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
  @GrpcMethod()
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
  @GrpcMethod()
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
  @GrpcMethod()
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
  // async upvoteThread(data: ThreadIdRequestBody, metadata: any) {
  //   const thread = await this.ThreadService.findOneById(data.thread_id);

  //   const threadId = thread._id;
  //   const updatedUpvoted = thread.upvoted + 1;

  //   const updatedThread = await this.ThreadService.updateThreadById(threadId, {
  //     upvoted: updatedUpvoted,
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
