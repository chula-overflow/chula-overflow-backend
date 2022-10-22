import { ObjectId } from 'mongoose';

export interface CourseBody {
  _id: ObjectId;
  course_id: string;
  course_name: string;
  course_codename: string;
  exam_ids?: string[];
}

export interface CourseCreateBody {
  course_id: string;
  course_name: string;
  course_codename: string;
}

export interface CourseUpdateBody {
  couse_id?: string;
  course_name?: string;
  course_codename?: string;
  exam_ids?: string[];
}
