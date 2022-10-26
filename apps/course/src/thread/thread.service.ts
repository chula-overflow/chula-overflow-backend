import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model, ObjectId } from 'mongoose';
import {
  ThreadAnswerUpdateBody,
  ThreadBody,
  ThreadCreateBody,
  ThreadUpdateBody,
} from './thread.interface';
import { Thread, ThreadDocument } from './thread.schema';

@Injectable()
export class ThreadService {
  constructor(
    @InjectModel(Thread.name)
    private ThreadModel: Model<ThreadDocument>,
  ) {}

  async createThread(threadBody: ThreadCreateBody): Promise<ThreadBody> {
    const newThread = new this.ThreadModel(threadBody);
    return newThread.save();
  }

  async generateTitle() {}

  async findByExamId(examId: ObjectId) {
    const threads = await this.ThreadModel.find({
      exam_id: examId,
    }).exec();
    return threads;
  }

  async findOneById(threadId: ObjectId) {
    const thread = await this.ThreadModel.findById(threadId);
    return thread;
  }

  async updateThreadById(threadId: ObjectId, threadBody: ThreadUpdateBody) {
    const updatedThread = this.ThreadModel.findByIdAndUpdate(
      threadId,
      { ...threadBody },
      { new: true },
    );
    return updatedThread;
  }

  async updateProblemByProblemId(problemId: string) {}

  async updateAnswerBodyByAnswerId(
    answerId: string,
    answerBody: ThreadAnswerUpdateBody,
  ) {
    const updatedThread = this.ThreadModel.updateOne(
      { 'answers.id': answerId },
      {
        $set: {
          'answer.$.body': answerBody.body,
        },
      },
    );
    return updatedThread;
  }

  async updateAnswerUpvoteByAnswerId(
    answerId: string,
    answerBody: ThreadAnswerUpdateBody,
  ) {
    const updatedThread = this.ThreadModel.updateOne(
      { 'answers.id': answerId },
      {
        $set: {
          'answer.$.upvoted': answerBody.upvoted,
        },
      },
    );
    return updatedThread;
  }

  async updateAnswerDownvoteByAnswerId(
    answerId: string,
    answerBody: ThreadAnswerUpdateBody,
  ) {
    const updatedThread = this.ThreadModel.updateOne(
      { 'answers.id': answerId },
      {
        $set: {
          'answer.$.downvoted': answerBody.downvoted,
        },
      },
    );
    return updatedThread;
  }
}
