import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type ThreadDocument = Thread & Document;

@Schema({ timestamps: true })
export class Thread {
  @Prop({ required: true })
  exam_id: string;

  @Prop({ required: true })
  course_id: string;

  @Prop({ required: true })
  upvoted: number;

  @Prop({ required: true })
  downvoted: number;

  @Prop({ required: true })
  problems: {
    id: string;
    title: string;
    body: string;
    uploaded_user: string;
    upvoted: number;
    downvoted: number;
  }[];

  @Prop({})
  answers: {
    id: string;
    body: string;
    upvoted: number;
    downvoted: number;
  }[];
}

export const ThreadSchema = SchemaFactory.createForClass(Thread);
