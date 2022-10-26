import { Controller, Post, Get, Res, Body, Param } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { IS_MICROSERVICE } from 'src/main';
import {
  CourseCreateBody,
  CourseRequestBody,
  CourseRequestUpdateBody,
} from './course.interface';
import { CourseService } from './course.service';

@Controller('course')
export class CourseController {
  constructor(private readonly CourseService: CourseService) {}

  @Post('/')
  @GrpcMethod('Course')
  async createCourse(
    @Res() response,
    @Body() reqBody: CourseCreateBody,
    data: CourseCreateBody,
    metadata: any,
  ) {
    const newCourse = await this.CourseService.create(
      IS_MICROSERVICE ? data : reqBody,
    );

    return IS_MICROSERVICE ? newCourse : response.status(201).json(newCourse);
  }

  @Get('/')
  @GrpcMethod('Course')
  async getAllCourses(@Res() response, metadata: any) {
    const courses = await this.CourseService.find();

    return IS_MICROSERVICE ? courses : response.status(200).json(courses);
  }

  @Get('/:courseId')
  @GrpcMethod('Course')
  async getCourseByCourseId(
    @Res() response,
    @Param('courseId') courseId,
    data: CourseRequestBody,
    metadata: any,
  ) {
    const course = await this.CourseService.findOneByCourseId(
      IS_MICROSERVICE ? data.course_id : courseId,
    );
    return IS_MICROSERVICE ? course : response.status(200).json(course);
  }

  @GrpcMethod('Course')
  async updateCourse(data: CourseRequestUpdateBody, metadata: any) {
    const updatedCourse = await this.CourseService.updateByCourseId(
      data.course_id,
      data.body,
    );
    return updatedCourse;
  }

  @GrpcMethod('Course')
  async deleteCourse(data: CourseRequestBody, metadata: any) {
    const deletedCourse = await this.CourseService.deleteByCourseId(
      data.course_id,
    );
    return deletedCourse;
  }
}
