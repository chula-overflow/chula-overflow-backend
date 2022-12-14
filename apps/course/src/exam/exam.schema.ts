import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type ExamDocument = Exam & Document;

@Schema()
export class Exam {
  @Prop({ required: true })
  course_id: string;

  @Prop({ required: true })
  year: number;

  @Prop({ required: true })
  semester: string;

  @Prop({ required: true })
  term: string;

  @Prop()
  thread_ids: string[];
}

export const ExamSchema = SchemaFactory.createForClass(Exam);
