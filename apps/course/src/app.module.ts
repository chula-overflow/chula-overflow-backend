import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MicrController } from './micr/micr.controller';
import { Course, CourseSchema } from './models/course.schema';
import { CourseController } from './course/course.controller';
import { CourseService } from './course/course.service';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MongooseModule.forRoot(process.env.MONGODB_URI),
    MongooseModule.forFeature([{ name: Course.name, schema: CourseSchema }]),
  ],
  controllers: [AppController, MicrController, CourseController],
  providers: [AppService, CourseService],
})
export class AppModule {}
