import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { Types } from 'mongoose';
import { ExamPropertyRequestBody } from 'src/exam/exam.interface';
import { ExamService } from 'src/exam/exam.service';
import {
  ThreadIdRequestBody,
  ThreadRequestAddAnswerBody,
  ThreadRequestCreateBody,
} from './thread.interface';
import { ThreadService } from './thread.service';

@Controller('thread')
export class ThreadController {
  constructor(
    private readonly ThreadService: ThreadService,
    private readonly ExamService: ExamService,
  ) {}

  @GrpcMethod()
  async createThread(data: ThreadRequestCreateBody, metadata: any) {
    const examId = await this.ExamService.findIdByCourseProperty(
      data.year,
      data.semester,
      data.term,
    );

    const createThreadBody = {
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
      answers: [
        data.answer && {
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

    return newThread;
  }

  @GrpcMethod()
  async getAllThreadsByExamProperty(
    data: ExamPropertyRequestBody,
    metadata: any,
  ) {
    const examId = await this.ExamService.findIdByCourseProperty(
      data.year,
      data.semester,
      data.term,
    );

    const threads = await this.ThreadService.findByExamId(examId);

    return threads;
  }

  @GrpcMethod()
  async getThreadById(data: ThreadIdRequestBody, metadata: any) {
    const thread = await this.ThreadService.findOneById(data.thread_id);

    return thread;
  }

  @GrpcMethod()
  async getThreadByTitle(data, metadata: any) {}

  @GrpcMethod()
  async addAnswerById(data: ThreadRequestAddAnswerBody, metadata: any) {
    const thread = await this.ThreadService.findOneById(data.thread_id);

    const threadId = thread._id;
    const answer = {
      body: data.body,
      upvoted: 0,
      downvoted: 0,
    };

    const updatedThread = await this.ThreadService.updateThreadById(threadId, {
      answers: [...thread.answers, answer],
    });

    return updatedThread;
  }

  @GrpcMethod()
  async upvoteThread(data: ThreadIdRequestBody, metadata: any) {
    const thread = await this.ThreadService.findOneById(data.thread_id);

    const threadId = thread._id;
    const updatedUpvoted = thread.upvoted + 1;

    const updatedThread = await this.ThreadService.updateThreadById(threadId, {
      upvoted: updatedUpvoted,
    });

    return updatedThread;
  }

  @GrpcMethod()
  async downvoteThread(data: ThreadIdRequestBody, metadata: any) {
    const thread = await this.ThreadService.findOneById(data.thread_id);

    const threadId = thread._id;
    const updatedDownvoted = thread.downvoted - 1;

    const updatedThread = await this.ThreadService.updateThreadById(threadId, {
      upvoted: updatedDownvoted,
    });

    return updatedThread;
  }
}
