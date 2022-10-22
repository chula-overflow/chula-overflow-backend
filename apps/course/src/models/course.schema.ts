import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type CourseDocument = Course & Document;

@Schema()
export class Course {
  @Prop({ required: true })
  courseName: string;

  @Prop({ required: true })
  courseCodename: string;

  @Prop({ required: true })
  courseId: string;

  @Prop({ required: true })
  exams: {
    term: string;
    year: number;
    threads: {
      title: string;
      description: string;
      problems: {
        description: string;
        answer: string;
        upvoted: number;
        downvoted: number;
      }[];
    }[];
    createdAt: Date;
  }[];
}

export const CourseSchema = SchemaFactory.createForClass(Course);
