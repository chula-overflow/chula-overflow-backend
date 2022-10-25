import { ObjectId } from 'mongoose';

export interface ThreadBody {
  _id: ObjectId;
  examId: stringd;
  courseId: string;
  upvoted: number;
  downvoted: number;
  problems: {
    title: string; // generated from nlp
    body: string;
    uploadedUser: string;
    upvoted: number;
    downvoted: number;
  }[];
  answers?: {
    body: string;
    upvoted: number;
    downvoted: number;
  }[];
}

export interface ThreadRequestCreateBody {
  courseId: string;
  year: number;
  semester: string;
  term: string;

  uploadedUser: string;

  question: string;
  answer?: string;
}

export interface ThreadCreateBody {
  examId: string;
  courseId: string;
  upvoted: number;
  downvoted: number;
  problems: {
    title: string;
    body: string;
    uploadedUser: string;
    upvoted: number;
    downvoted: number;
  }[];
  answers?: {
    body: string;
    upvoted: number;
    downvoted: number;
  }[];
}

export interface ThreadUpdateBody {
  examId?: string;
  courseId?: string;
  upvoted?: number;
  downvoted?: number;
  problems?: {
    title?: string;
    body?: string;
    uploadedUser?: string;
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
  threadId: ObjectId;
}
