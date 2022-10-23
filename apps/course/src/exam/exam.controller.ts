import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { CourseService } from 'src/course/course.service';
import {
  ExamCourseIdRequestBody,
  ExamCreateBody,
  ExamPropertyRequestBody,
  ExamRequestUpdateBody,
} from './exam.interface';
import { ExamService } from './exam.service';

@Controller('exam')
export class ExamController {
  constructor(
    private readonly ExamService: ExamService,
    private readonly CourseService: CourseService,
  ) {}

  @GrpcMethod()
  async createExam(data: ExamCreateBody, metadata: any) {
    const newExam = await this.ExamService.create(data);
    const examId = newExam._id;
    const courseId = newExam.course_id;

    // add exam_id to course document
    await this.CourseService.addExamId(examId, courseId);

    return newExam;
  }

  @GrpcMethod()
  async getAllExams(metadata: any) {
    const exams = await this.ExamService.find();
    return exams;
  }

  @GrpcMethod()
  async getAllExamsByCourseId(data: ExamCourseIdRequestBody, metadata: any) {
    const exams = await this.ExamService.findByCourseId(data.course_id);
    return exams;
  }

  @GrpcMethod()
  async getExamByCourseProperty(data: ExamPropertyRequestBody, metadata: any) {
    const examId = await this.ExamService.findIdByCourseProperty(
      data.year,
      data.semester,
      data.term,
    );

    const exam = await this.ExamService.findOneById(examId);

    return exam;
  }

  @GrpcMethod()
  async updateExamByCourseProperty(data: ExamRequestUpdateBody, metadata: any) {
    const examId = await this.ExamService.findIdByCourseProperty(
      data.year,
      data.semester,
      data.term,
    );

    const updatedExam = await this.ExamService.updateById(examId, data.body);

    return updatedExam;
  }

  @GrpcMethod()
  async deleteExamByCourseProperty(
    data: ExamPropertyRequestBody,
    metadata: any,
  ) {
    const examId = await this.ExamService.findIdByCourseProperty(
      data.year,
      data.semester,
      data.term,
    );

    const deletedExam = await this.ExamService.deleteById(examId);

    return deletedExam;
  }
}
