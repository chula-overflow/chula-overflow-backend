import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import {
  CourseCreateBody,
  CourseRequestBody,
  CourseRequestUpdateBody,
  CourseUpdateBody,
} from './course.interface';
import { CourseService } from './course.service';

@Controller('course')
export class CourseController {
  constructor(private readonly CourseService: CourseService) {}

  @GrpcMethod()
  async createCourse(data: CourseCreateBody, metadata: any) {
    const newCourse = await this.CourseService.create(data);
    return newCourse;
  }

  @GrpcMethod()
  async getAllCourses(metadata: any) {
    const courses = await this.CourseService.find();
    return courses;
  }

  @GrpcMethod()
  async getCourseByCourseId(data: CourseRequestBody, metadata: any) {
    const course = await this.CourseService.findOneByCourseId(data.course_id);
    return course;
  }

  @GrpcMethod()
  async updateCourse(data: CourseRequestUpdateBody, metadata: any) {
    const updatedCourse = await this.CourseService.updateByCourseId(
      data.course_id,
      data.body,
    );
    return updatedCourse;
  }

  @GrpcMethod()
  async deleteCourse(data: CourseRequestBody, metadata: any) {
    const deletedCourse = await this.CourseService.deleteByCourseId(
      data.course_id,
    );
    return deletedCourse;
  }
}
