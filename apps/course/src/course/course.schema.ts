import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type CourseDocument = Course & Document;

@Schema()
export class Course {
  @Prop({ required: true })
  course_id: string;

  @Prop({ required: true })
  course_name: string;

  @Prop({ required: true })
  course_codename: string;

  @Prop()
  exam_ids: string[];
}

export const CourseSchema = SchemaFactory.createForClass(Course);
