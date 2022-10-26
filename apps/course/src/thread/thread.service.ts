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
  ) {}

  async createThread(threadBody: ThreadCreateBody): Promise<ThreadBody> {
    const newThread = new this.ThreadModel(threadBody);
    return newThread.save();
  }

  async find(threadProperty: {
    exam_id: string;
    course_id: string;
  }): Promise<ThreadBody[]> {
    const threads = this.ThreadModel.find({
      ...threadProperty,
    }).exec();
    return threads;
  }

  async findAll() {
    const threads = this.ThreadModel.find({}).exec();
    return threads;
  }

  async findByExamId(examId: string) {
    const thread = await this.ThreadModel.findOne({
      exam_id: examId,
    }).exec();
    return thread;
  }

  async findOneById(threadId: ObjectId) {
    const thread = await this.ThreadModel.findById(threadId);
    return thread;
  }

  async findProblem(problemId: string) {
    const thread = await this.ThreadModel.findOne({
      'problems.id': problemId,
    }).exec();

    const problems = thread.problems;
    const problem = problems.find((problem) => problem.id == problemId);

    return problem;
  }

  async findAnswer(answerId: string) {
    const thread = await this.ThreadModel.findOne({
      'answers.id': answerId,
    }).exec();

    const answers = thread.answers;
    const answer = answers.find((answer) => answer.id == answerId);

    return answer;
  }

  async updateVoted(
    type: {
      problem?: boolean;
      answer?: boolean;
      upvote?: boolean;
      downvote?: boolean;
    },
    id: string,
    voted: number,
  ) {
    const updatedThread = await this.ThreadModel.updateOne(
      type.problem ? { 'problems.id': id } : { 'answers.id': id },
      {
        $set: type.problem
          ? type.upvote
            ? {
                'problems.$.upvoted': voted,
              }
            : {
                'problems.$.downvoted': voted,
              }
          : type.upvote
          ? {
              'answers.$.upvoted': voted,
            }
          : {
              'answers.$.downvoted': voted,
            },
      },
    );
    return updatedThread;
  }

  async updateThreadById(threadId: ObjectId, threadBody: ThreadUpdateBody) {
    const updatedThread = await this.ThreadModel.findByIdAndUpdate(
      threadId,
      { ...threadBody },
      { new: true },
    );

    return updatedThread;
  }
}
