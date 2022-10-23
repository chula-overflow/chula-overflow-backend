import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ExamService } from 'src/exam/exam.service';
import { ThreadRequestCreateBody } from './thread.interface';
import { ThreadService } from './thread.service';

@Controller('thread')
export class ThreadController {
  constructor(
    private readonly ThreadService: ThreadService,
    private readonly ExamService: ExamService,
  ) {}

  @GrpcMethod()
  async createThread(data: ThreadRequestCreateBody, metadata: any) {
    const examId = await this.ExamService.findIdByCourseProperty(
      data.year,
      data.semester,
      data.term,
    );

    const createThreadBody = {
      exam_id: examId,
      course_id: data.course_id,
      upvoted: 0,
      downvoted: 0,
      problems: [
        {
          title: 'generated from nlp',
          body: data.question,
          uploaded_user: data.uploaded_user,
          upvoted: 0,
          downvoted: 0,
        },
      ],
      answers: [
        data.answer && {
          body: data.answer,
          upvoted: 0,
          downvoted: 0,
        },
      ],
    };

    const newThread = await this.ThreadService.createThread(createThreadBody);

    const threadId = newThread._id;

    // add thread_id to exam document
    await this.ExamService.addThreadId(threadId, examId);

    return newThread;
  }
}
