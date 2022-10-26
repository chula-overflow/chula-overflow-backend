import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MicrController } from './micr/micr.controller';
import { Course, CourseSchema } from './course/course.schema';
import { CourseController } from './course/course.controller';
import { CourseService } from './course/course.service';
import { ExamController } from './exam/exam.controller';
import { ExamService } from './exam/exam.service';
import { ThreadController } from './thread/thread.controller';
import { ThreadService } from './thread/thread.service';
import { Exam, ExamSchema } from './exam/exam.schema';
import { Thread, ThreadSchema } from './thread/thread.schema';
import { EmbededController } from './embeded/embeded.controller';
import { Embeded, EmbededSchema } from './embeded/embeded.schema';
import { EmbededService } from './embeded/embeded.service';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MongooseModule.forRoot(process.env.MONGODB_URI),
    MongooseModule.forFeature([
      { name: Course.name, schema: CourseSchema },
      { name: Exam.name, schema: ExamSchema },
      { name: Thread.name, schema: ThreadSchema },
      { name: Embeded.name, schema: EmbededSchema },
    ]),
  ],
  controllers: [
    AppController,
    MicrController,
    CourseController,
    ExamController,
    ThreadController,
    EmbededController,
  ],
  providers: [AppService, CourseService, ExamService, ThreadService, EmbededService],
})
export class AppModule {}
