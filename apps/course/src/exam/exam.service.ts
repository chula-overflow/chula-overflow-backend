import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { CourseDocument } from 'src/course/course.schema';
import { CourseService } from 'src/course/course.service';
import { Exam, ExamDocument } from 'src/exam/exam.schema';
import { ExamBody } from './exam.interface';

@Injectable()
export class ExamService {
  constructor(
    @InjectModel(Exam.name)
    private ExamModel: Model<ExamDocument>,

    @InjectModel(Exam.name)
    private CourseModel: Model<CourseDocument>,

    private readonly CourseService: CourseService,
  ) {}

  // create new exam and update that course to has exam id
  async create(examData: ExamBody): Promise<ExamBody> {
    const newExam = new this.ExamModel(examData);
    const examId = newExam._id;
    const courseId = newExam.course_id;

    // add exam id to course
    await this.CourseService.addExamId(examId, courseId);

    return newExam.save();
  }

  // find all exams
  async find() {
    const exams = await this.ExamModel.find({}).exec();
    return exams;
  }

  async findOneByCourseId(courseId: string) {
    const exam = await this.ExamModel.findOne({
      course_id: courseId,
    }).exec();
    return exam;
  }

  async findOneByProperty(year: number, semester: string, term: string) {
    const exam = await this.ExamModel.findOne({
      year: year,
      semester: semester,
      term: term,
    }).exec();

    return exam;
  }

  async addThreadId(threadId: string) {}

  async update() {}

  async delete() {}
}
