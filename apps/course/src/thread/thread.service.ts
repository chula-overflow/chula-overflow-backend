import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { CourseService } from 'src/course/course.service';
import { Exam, ExamDocument } from 'src/exam/exam.schema';
import { ExamService } from 'src/exam/exam.service';
import { AddProblemBody, CreateThreadBody } from './thread.interface';

@Injectable()
export class ThreadService {
  constructor(
    @InjectModel(Exam.name)
    private ThreadModel: Model<ExamDocument>,

    private readonly CourseService: CourseService,
    private readonly ExamService: ExamService,
  ) {}

  async createThread(threadBody: CreateThreadBody) {
    const examId = (
      await this.ExamService.findOneByProperty(
        threadBody.year,
        threadBody.semester,
        threadBody.term,
      )
    )._id;
    const threadData = {
      course_id: threadBody.course_id,
      exam_id: examId,
      upvoted: 0,
      downvoted: 0,
      problems: [],
      answers: [],
    };
    const newThread = new this.ThreadModel(threadData);

    return newThread.save();
  }

  async findOneByExamId(examId: string) {
    const thread = await this.ThreadModel.findOne({
      exam_id: examId,
    }).exec();
    return thread;
  }

  async addProblem(problemData: AddProblemBody) {
    const thread = this.findOneByExamId(problemData.exam_id);
  }
}
