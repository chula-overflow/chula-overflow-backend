import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model, ObjectId } from 'mongoose';
import { Course, CourseDocument } from 'src/course/course.schema';
import {
  CourseBody,
  CourseCreateBody,
  CourseUpdateBody,
} from './course.interface';

@Injectable()
export class CourseService {
  constructor(
    @InjectModel(Course.name)
    private CourseModel: Model<CourseDocument>,
  ) {}

  async create(courseData: CourseCreateBody): Promise<CourseBody> {
    const newCourse = new this.CourseModel(courseData);
    return newCourse.save(); // return created course
  }

  async find(): Promise<CourseBody[]> {
    const courses = await this.CourseModel.find({}).exec();
    return courses; // return all courses with all properties
  }

  async findOneByCourseId(courseId: string): Promise<CourseBody> {
    const course = await this.CourseModel.findOne({
      course_id: courseId,
    }).exec();
    return course; // return each course with all properties
  }

  async addExamId(examId: ObjectId, courseId: string): Promise<CourseBody> {
    const course = await this.findOneByCourseId(courseId);

    const courseDocumentId = course._id;
    const examIds = [...course.exam_ids, examId];

    const updatedCourse = await this.CourseModel.findByIdAndUpdate(
      courseDocumentId,
      { exam_ids: examIds },
      { new: true },
    );

    return updatedCourse;
  }

  async updateByCourseId(
    courseId: string,
    updateBody: CourseUpdateBody,
  ): Promise<CourseBody> {
    const updatedCourse = await this.CourseModel.findOneAndUpdate(
      { course_id: courseId },
      { ...updateBody },
      { new: true },
    );

    return updatedCourse;
  }

  async deleteByCourseId(courseId: string): Promise<CourseBody> {
    const deletedCourse = await this.CourseModel.findOneAndRemove({
      course_id: courseId,
    });
    return deletedCourse;
  }
}
