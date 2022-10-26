import { ObjectId } from 'mongoose';

export interface ExamBody {
  _id: ObjectId;
  course_id: string;
  year: number; // 2022, 2021, 2020, ...
  semester: string; // S1, S2
  term: string; // Mid | Final
  thread_ids?: string[];
}

export interface ExamCreateBody {
  course_id: string;
  year: number; // 2022, 2021, 2020, ...
  semester: string; // S1, S2
  term: string; // Mid | Final
}

export interface ExamUpdateBody {
  course_id?: string;
  year?: number; // 2022, 2021, 2020, ...
  semester?: string; // S1, S2
  term?: string; // Mid | Final
  thread_ids?: string[];
}

export interface ExamCourseIdRequestBody {
  course_id: string;
}

export interface ExamPropertyRequestBody {
  course_id?: string;
  year: number; // 2022, 2021, 2020, ...
  semester: string; // S1, S2
  term: string; // Mid | Final
}

export interface ExamRequestUpdateBody {
  year: number; // 2022, 2021, 2020, ...
  semester: string; // S1, S2
  term: string; // Mid | Final
  body: ExamUpdateBody;
}
