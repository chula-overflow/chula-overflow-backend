import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Exam, ExamDocument } from 'src/exam/exam.schema';
import { ThreadCreateBody } from './thread.interface';

@Injectable()
export class ThreadService {
  constructor(
    @InjectModel(Exam.name)
    private ThreadModel: Model<ExamDocument>,
  ) {}

  async createThread(threadBody: ThreadCreateBody) {
    const newThread = new this.ThreadModel(threadBody);

    return newThread.save();
  }

  async findOneByExamId(examId: string) {
    const thread = await this.ThreadModel.findOne({
      exam_id: examId,
    }).exec();
    return thread;
  }

  async addProblem(problemData) {
    const thread = this.findOneByExamId(problemData.exam_id);
  }
}
