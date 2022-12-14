import { ObjectId } from 'mongoose';

export interface ThreadBody {
  _id: ObjectId;
  exam_id: string;
  course_id: string;
  upvoted: number;
  downvoted: number;
  problems: {
    id: string;
    title: string; // generated from nlp
    body: string;
    uploaded_user: string;
    upvoted: number;
    downvoted: number;
  }[];
  answers?: {
    id: string;
    body: string;
    upvoted: number;
    downvoted: number;
  }[];
}

export interface ThreadRequestCreateBody {
  course_id: string;
  year: number;
  semester: string;
  term: string;

  uploaded_user: string;

  question: string;
  question_title: string;
  answer?: string;
}

export interface ThreadCreateBody {
  exam_id: string;
  course_id: string;
  upvoted: number;
  downvoted: number;
  problems: {
    id: string;
    title: string;
    body: string;
    uploaded_user: string;
    upvoted: number;
    downvoted: number;
  }[];
  answers?: {
    id: string;
    body: string;
    upvoted: number;
    downvoted: number;
  }[];
}

export interface ThreadRequestBody {
  course_id: string;
  exam_id?: string;
  year: number;
  semester: string;
  term: string;
}

export interface ThreadUpdateBody {
  exam_id?: string;
  course_id?: string;
  upvoted?: number;
  downvoted?: number;
  problems?: {
    title?: string;
    body?: string;
    uploaded_user?: string;
    upvoted?: number;
    downvoted?: number;
  }[];
  answers?: {
    body?: string;
    upvoted?: number;
    downvoted?: number;
  }[];
}

export interface ThreadIdRequestBody {
  thread_id: ObjectId;
}

export interface ThreadAnswerUpdateBody {
  body?: string;
  upvoted?: number;
  downvoted?: number;
}

export interface ThreadRequestAddAnswerBody {
  thread_id: ObjectId;
  body: string;
}
