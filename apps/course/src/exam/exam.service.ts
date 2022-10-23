import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model, ObjectId } from 'mongoose';
import { Exam, ExamDocument } from 'src/exam/exam.schema';
import { ExamBody, ExamCreateBody, ExamUpdateBody } from './exam.interface';

@Injectable()
export class ExamService {
  constructor(
    @InjectModel(Exam.name)
    private ExamModel: Model<ExamDocument>,
  ) {}

  // create new exam and update that course to has exam id
  async create(examData: ExamCreateBody): Promise<ExamBody> {
    const newExam = new this.ExamModel(examData);
    return newExam.save();
  }

  // find all exams
  async find(): Promise<ExamBody[]> {
    const exams = await this.ExamModel.find({}).exec();
    return exams;
  }

  async findByCourseId(courseId: string): Promise<ExamBody[]> {
    const exam = await this.ExamModel.find({
      course_id: courseId,
    }).exec();
    return exam;
  }

  async findIdByCourseProperty(
    year: number,
    semester: string,
    term: string,
  ): Promise<ObjectId> {
    const exam = await this.ExamModel.findOne(
      {
        year: year,
        semester: semester,
        term: term,
      },
      '_id',
    ).exec();

    const exam_id = exam._id;

    return exam_id;
  }

  async findOneById(examId: ObjectId): Promise<ExamBody> {
    const exam = await this.ExamModel.findById(examId).exec();

    return exam;
  }

  async addThreadId(threadId: ObjectId, examId: ObjectId): Promise<ExamBody> {
    const exam = await this.findOneById(examId);

    const examDocumentId = exam._id;
    const threadIds = [...exam.thread_ids, String(threadId)];

    const updatedExam = await this.ExamModel.findByIdAndUpdate(
      examDocumentId,
      { thread_ids: threadIds },
      { new: true },
    );

    return updatedExam;
  }

  async updateById(
    examId: ObjectId,
    updateBody: ExamUpdateBody,
  ): Promise<ExamBody> {
    const updatedExam = await this.ExamModel.findByIdAndUpdate(
      examId,
      { ...updateBody },
      { new: true },
    );

    return updatedExam;
  }

  async deleteById(examId: ObjectId): Promise<ExamBody> {
    const deletedExam = await this.ExamModel.findByIdAndRemove(examId);
    return deletedExam;
  }
}
