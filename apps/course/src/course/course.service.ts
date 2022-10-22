import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Course, CourseDocument } from 'src/course/course.schema';
import { CourseInterface } from './course.interface';

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
    const courses = await this.CourseModel.find({}).exec();
    return courses; // return all courses with all properties
  }

  async findOneByCourseId(courseId: string) {
    const course = await this.CourseModel.findOne({
      courseId: courseId,
    }).exec();
    return course; // return each course with all properties
  }

  async addExamId(examId: string, courseId: string) {
    const course = await this.CourseModel.findOne({
      courseId: courseId,
    }).exec();
    const courseDocumentId = course._id;
    const examsId = [...course.exams_id, examId];
    const updatedCourse = await this.CourseModel.findByIdAndUpdate(
      courseDocumentId,
      { exams_id: examsId },
      { new: true },
    );

    return updatedCourse;
  }

  async update(courseId: string) {}

  async delete(courseId: string) {}
}
