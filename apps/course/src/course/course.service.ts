import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Course, CourseDocument } from 'src/models/course.schema';
import { CourseInterface } from './course.interface';

@Injectable()
export class CourseService {
  constructor(
    @InjectModel(Course.name)
    private CourseModel: Model<CourseDocument>,
  ) {}

  async create(courseData: CourseInterface): Promise<CourseInterface> {
    const newCourse = new this.CourseModel(courseData);
    return newCourse.save();
  }

  async find(): Promise<CourseInterface[]> {
    const courses = this.CourseModel.find().exec();
    return courses;
  }
}
