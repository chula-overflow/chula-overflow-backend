import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Course, CourseDocument } from 'src/models/course.schema';

@Injectable()
export class ExamService {
  constructor(
    @InjectModel(Course.name)
    private CourseModel: Model<CourseDocument>,
  ) {}

  async create() {}

  async find() {}
}
