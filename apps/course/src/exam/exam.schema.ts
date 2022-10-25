import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type ExamDocument = Exam & Document;

@Schema()
export class Exam {
  @Prop({ required: true })
  courseId: string;

  @Prop({ required: true })
  year: number;

  @Prop({ required: true })
  semester: string;

  @Prop({ required: true })
  term: string;

  @Prop()
  threadIds: string[];
}

export const ExamSchema = SchemaFactory.createForClass(Exam);
