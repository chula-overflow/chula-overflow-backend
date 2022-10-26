import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type EmbededDocument = Embeded & Document;

@Schema({ timestamps: true })
export class Embeded {
  @Prop({ required: true })
  problem_id: string;

  @Prop({ required: true })
  vector: number[];
}

export const EmbededSchema = SchemaFactory.createForClass(Embeded);
