import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { ThreadBody, ThreadCreateBody } from './thread.interface';
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

  async findByExamId(examId: string) {
    const threads = await this.ThreadModel.find({
      exam_id: examId,
    }).exec();
    return threads;
  }
}
