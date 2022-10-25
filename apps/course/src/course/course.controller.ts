import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import {
  CourseCreateBody,
  CourseRequestBody,
  CourseRequestUpdateBody,
} from './course.interface';
import { CourseService } from './course.service';

@Controller('course')
export class CourseController {
  constructor(private readonly CourseService: CourseService) { }

  @GrpcMethod('Course')
  async createCourse(data: CourseCreateBody, metadata: any) {
    const newCourse = await this.CourseService.create(data);
    return newCourse;
  }

  @GrpcMethod('Course')
  async getAllCourses(data: {}, metadata: any) {
    const courses = await this.CourseService.find();
    return {
      messages: courses,
    };
  }

  @GrpcMethod('Course')
  async getCourseByCourseId(data: CourseRequestBody, metadata: any) {
    const course = await this.CourseService.findOneByCourseId(data.courseId);
    return course;
  }

  @GrpcMethod('Course')
  async updateCourse(data: CourseRequestUpdateBody, metadata: any) {
    const updatedCourse = await this.CourseService.updateByCourseId(
      data.courseId,
      data.body,
    );
    return updatedCourse;
  }

  @GrpcMethod('Course')
  async deleteCourse(data: CourseRequestBody, metadata: any) {
    const deletedCourse = await this.CourseService.deleteByCourseId(
      data.courseId,
    );
    return deletedCourse;
  }
}
