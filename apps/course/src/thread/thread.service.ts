import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model, ObjectId } from 'mongoose';
import {
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
  ) { }

  async createThread(threadBody: ThreadCreateBody): Promise<ThreadBody> {
    const newThread = new this.ThreadModel(threadBody);
    return newThread.save();
  }

  async findByExamId(examId: ObjectId) {
    const threads = await this.ThreadModel.find({
      examId: examId,
    }).exec();
    return threads;
  }

  async findOneById(threadId: ObjectId) {
    const thread = await this.ThreadModel.findById(threadId);
    return thread;
  }

  async updateById(threadId: ObjectId, threadBody: ThreadUpdateBody) {
    const updatedThread = this.ThreadModel.findByIdAndUpdate(
      threadId,
      { ...threadBody },
      { new: true },
    );
    return updatedThread;
  }
}
