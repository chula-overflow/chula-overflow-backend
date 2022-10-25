import { ObjectId } from 'mongoose';

export interface CourseBody {
  _id: ObjectId;
  courseId: string;
  courseName: string;
  courseCodename: string;
  examIds?: string[];
}

export interface CourseCreateBody {
  courseId: string;
  courseName: string;
  courseCodename: string;
}

export interface CourseUpdateBody {
  couseId?: string;
  courseName?: string;
  courseCodename?: string;
  examIds?: string[];
}

export interface CourseRequestBody {
  courseId: string;
}

export interface CourseRequestUpdateBody {
  courseId: string;
  body: CourseUpdateBody;
}
