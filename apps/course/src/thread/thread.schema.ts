import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document, ObjectId } from 'mongoose';

export type ThreadDocument = Thread & Document;

@Schema({ timestamps: true })
export class Thread {
  @Prop({ required: true })
  exam_id: string;

  @Prop({ required: true })
  course_id: string;

  @Prop({ required: true })
  title: string;

  @Prop({ required: true })
  upvoted: number;

  @Prop({ required: true })
  downvoted: number;

  @Prop({ required: true })
  problems: {
    title: string;
    body: string;
    uploaded_user: string;
    upvoted: number;
    downvoted: number;
  }[];

  @Prop({})
  answers: {
    body: string;
    upvoted: number;
    downvoted: number;
  }[];
}

export const ThreadSchema = SchemaFactory.createForClass(Thread);
