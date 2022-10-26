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

  async find(threadProperty: {
    exam_id: string;
    course_id: string;
  }): Promise<ThreadBody[]> {
    const threads = this.ThreadModel.find({
      ...threadProperty,
    }).exec();
    return threads;
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
    const updatedThread = this.ThreadModel.findByIdAndUpdate(
      threadId,
      { ...threadBody },
      { new: true },
    );
    return updatedThread;
  }

  // async generateTitle() {}

  // async findByExamId(examId: ObjectId) {
  //   const threads = await this.ThreadModel.find({
  //     exam_id: examId,
  //   }).exec();
  //   return threads;
  // }

  // async findOneById(threadId: ObjectId) {
  //   const thread = await this.ThreadModel.findById(threadId);
  //   return thread;
  // }

  // async updateProblemByProblemId(problemId: string) {}

  // async updateAnswerBodyByAnswerId(
  //   answerId: string,
  //   answerBody: ThreadAnswerUpdateBody,
  // ) {
  //   const updatedThread = this.ThreadModel.updateOne(
  //     { 'answers.id': answerId },
  //     {
  //       $set: {
  //         'answer.$.body': answerBody.body,
  //       },
  //     },
  //   );
  //   return updatedThread;
  // }

  // async updateAnswerUpvoteByAnswerId(
  //   answerId: string,
  //   answerBody: ThreadAnswerUpdateBody,
  // ) {
  //   const updatedThread = this.ThreadModel.updateOne(
  //     { 'answers.id': answerId },
  //     {
  //       $set: {
  //         'answer.$.upvoted': answerBody.upvoted,
  //       },
  //     },
  //   );
  //   return updatedThread;
  // }

  // async updateAnswerDownvoteByAnswerId(
  //   answerId: string,
  //   answerBody: ThreadAnswerUpdateBody,
  // ) {
  //   const updatedThread = this.ThreadModel.updateOne(
  //     { 'answers.id': answerId },
  //     {
  //       $set: {
  //         'answer.$.downvoted': answerBody.downvoted,
  //       },
  //     },
  //   );
  //   return updatedThread;
  // }
}
