import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Course, CourseDocument } from 'src/models/course.schema';
import { CourseInterface, CoursePropertiesInterface } from './course.interface';

@Injectable()
export class CourseService {
  constructor(
    @InjectModel(Course.name)
    private CourseModel: Model<CourseDocument>,
  ) {}

  async create(courseData: CourseInterface): Promise<CourseInterface> {
    const newCourse = new this.CourseModel(courseData);
    return newCourse.save(); // return created course
  }

  async find(): Promise<CourseInterface[]> {
    const courses = this.CourseModel.find({}).exec();
    return courses; // return all courses with all properties
  }

  async findProperties(): Promise<CoursePropertiesInterface[]> {
    const courses = this.CourseModel.find(
      {},
      'courseName courseCodename courseId',
    ).exec();
    return courses; // return all courses with course's property
  }

  async findOneByCourseId(courseId: string) {
    const course = this.CourseModel.findOne({ courseId: courseId }).exec();
    return course; // return each course with all properties
  }

  async update(courseId: string) {}

  async delete(courseId: string) {}
}
