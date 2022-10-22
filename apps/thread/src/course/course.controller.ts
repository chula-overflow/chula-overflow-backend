import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { CourseInterface } from './course.interface';
import { CourseService } from './course.service';

@Controller('course')
export class CourseController {
  constructor(private readonly CourseService: CourseService) {}

  @GrpcMethod()
  async createCourse(data: CourseInterface, metadata: any) {
    const newCourse = await this.CourseService.create(data);
    return newCourse;
  }

  @GrpcMethod()
  async findAllCourses(metadata: any) {
    const courses = await this.CourseService.find();
    return courses;
  }
}