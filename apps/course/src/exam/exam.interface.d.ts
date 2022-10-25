import { ObjectId } from 'mongoose';

export interface ExamBody {
  _id: ObjectId;
  courseId: string;
  year: number; // 2022, 2021, 2020, ...
  semester: string; // S1, S2
  term: string; // Mid | Final
  threadIds?: string[];
}

export interface ExamCreateBody {
  courseId: string;
  year: number; // 2022, 2021, 2020, ...
  semester: string; // S1, S2
  term: string; // Mid | Final
}

export interface ExamUpdateBody {
  courseId?: string;
  year?: number; // 2022, 2021, 2020, ...
  semester?: string; // S1, S2
  term?: string; // Mid | Final
  threadIds?: string[];
}

export interface ExamCourseIdRequestBody {
  courseId: string;
}

export interface ExamPropertyRequestBody {
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
