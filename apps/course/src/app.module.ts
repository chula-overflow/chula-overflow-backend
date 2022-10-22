import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MicrController } from './micr/micr.controller';
import { Course, CourseSchema } from './models/course.schema';
import { CourseController } from './course/course.controller';
import { CourseService } from './course/course.service';
import { ExamController } from './exam/exam.controller';
import { ExamService } from './exam/exam.service';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MongooseModule.forRoot(process.env.MONGODB_URI),
    MongooseModule.forFeature([{ name: Course.name, schema: CourseSchema }]),
  ],
  controllers: [AppController, MicrController, CourseController, ExamController],
  providers: [AppService, CourseService, ExamService],
})
export class AppModule {}
