import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { CourseDocument } from 'src/course/course.schema';
import { CourseService } from 'src/course/course.service';
import { Exam, ExamDocument } from 'src/exam/exam.schema';

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
  async create(examData) {
    const newExam = new this.ExamModel(examData);
    const examId = newExam._id;
    const courseId = newExam.course_id;

    await this.CourseService.addExamId(examId, courseId);
    return newExam;
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

  async findByProperty(year: number, semester: string, term: string) {
    const exam = await this.ExamModel.find({
      year: year,
      semester: semester,
      term: term,
    }).exec();

    return exam;
  }

  async update() {}

  async delete() {}
}
