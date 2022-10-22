import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ExamBody } from './exam.interface';
import { ExamService } from './exam.service';

@Controller('exam')
export class ExamController {
  constructor(private readonly ExamService: ExamService) {}

  @GrpcMethod()
  async createExam(data: ExamBody, metadata: any) {
    const newExam = await this.ExamService.create(data);
    return newExam;
  }

  @GrpcMethod()
  async getAllExams(metadata: any) {
    const exams = await this.ExamService.find();
    return exams;
  }

  @GrpcMethod()
  async getExamByCourseId(data: { courseId: string }, metadata: any) {
    const exam = await this.ExamService.findOneByCourseId(data.courseId);
    return exam;
  }

  @GrpcMethod()
  async getExamByProperty(
    data: { year: number; semester: string; term: string },
    metadata: any,
  ) {
    const exam = await this.ExamService.findOneByProperty(
      data.year,
      data.semester,
      data.term,
    );
    return exam;
  }
}
