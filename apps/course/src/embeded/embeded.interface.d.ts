import { ObjectId } from 'mongoose';

export interface EmbededBody {
  _id: ObjectId;
  problem_id: string;
  vector: number[];
}

export interface EmbededCreateBody {
  problem_id: string;
  vector: number[];
}
