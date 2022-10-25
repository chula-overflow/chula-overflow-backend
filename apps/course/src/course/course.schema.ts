import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type CourseDocument = Course & Document;

@Schema()
export class Course {
  @Prop({ required: true })
  courseId: string;

  @Prop({ required: true })
  courseName: string;

  @Prop({ required: true })
  courseCodename: string;

  @Prop()
  examIds: string[];
}

export const CourseSchema = SchemaFactory.createForClass(Course);
