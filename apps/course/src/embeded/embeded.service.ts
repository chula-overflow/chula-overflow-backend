import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { EmbededBody, EmbededCreateBody } from './embeded.interface';
import { Embeded, EmbededDocument } from './embeded.schema';

@Injectable()
export class EmbededService {
  constructor(
    @InjectModel(Embeded.name)
    private EmbededModel: Model<EmbededDocument>,
  ) {}

  async create(embededBody: EmbededCreateBody): Promise<EmbededBody> {
    const newEmbeded = new this.EmbededModel(embededBody);
    return newEmbeded.save();
  }

  async findByProblemId(problemId: string) {
    const embeded = await this.EmbededModel.findOne({
      problem_id: problemId,
    }).exec();
    return embeded;
  }
}
