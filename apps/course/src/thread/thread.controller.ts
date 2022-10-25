import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ObjectId } from 'mongoose';
import { ExamPropertyRequestBody } from 'src/exam/exam.interface';
import { ExamService } from 'src/exam/exam.service';
import {
  ThreadIdRequestBody,
  ThreadRequestCreateBody,
} from './thread.interface';
import { ThreadService } from './thread.service';

@Controller('thread')
export class ThreadController {
  constructor(
    private readonly ThreadService: ThreadService,
    private readonly ExamService: ExamService,
  ) { }

  @GrpcMethod('Thread')
  async createThread(data: ThreadRequestCreateBody, metadata: any) {
    console.log(data);

    const examId = await this.ExamService.findIdByCourseProperty(
      data.year,
      data.semester,
      data.term,
    );

    const createThreadBody = {
      examId: String(examId),
      courseId: data.courseId,
      upvoted: 0,
      downvoted: 0,
      problems: [
        {
          title: 'generated from nlp',
          body: data.question,
          uploadedUser: data.uploadedUser,
          upvoted: 0,
          downvoted: 0,
        },
      ],
      answers: [
        data.answer && {
          body: data.answer,
          upvoted: 0,
          downvoted: 0,
        },
      ],
    };

    const newThread = await this.ThreadService.createThread(createThreadBody);

    const threadId = newThread._id;

    // add threadId to exam document
    await this.ExamService.addThreadId(threadId, examId);

    return newThread;
  }

  @GrpcMethod('Thread')
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

    return {
      messages: threads,
    };
  }

  @GrpcMethod('Thread')
  async getThreadById(data: ThreadIdRequestBody, metadata: any) {
    const thread = await this.ThreadService.findOneById(data.threadId);

    return thread;
  }

  @GrpcMethod('Thread')
  async getThreadByTitle(data, metadata: any) { }

  @GrpcMethod('Thread')
  async upvoteThread(data: ThreadIdRequestBody, metadata: any) {
    const thread = await this.ThreadService.findOneById(data.threadId);

    const threadId = thread._id;
    const updatedUpvoted = thread.upvoted + 1;

    const updatedThread = await this.ThreadService.updateById(threadId, {
      upvoted: updatedUpvoted,
    });

    return updatedThread;
  }

  @GrpcMethod('Thread')
  async downvoteThread(data: ThreadIdRequestBody, metadata: any) {
    const thread = await this.ThreadService.findOneById(data.threadId);

    const threadId = thread._id;
    const updatedDownvoted = thread.downvoted - 1;

    const updatedThread = await this.ThreadService.updateById(threadId, {
      upvoted: updatedDownvoted,
    });

    return updatedThread;
  }
}
