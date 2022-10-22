import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { CourseCreateBody, GetCourseByCourseIdBody } from './course.interface';
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
  async getCourseByCourseId(data: GetCourseByCourseIdBody, metadata: any) {
    const course = await this.CourseService.findOneByCourseId(data.course_id);
    return course;
  }
}
