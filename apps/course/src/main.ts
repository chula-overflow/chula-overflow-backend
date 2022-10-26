import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { NestExpressApplication } from '@nestjs/platform-express';
import { AppModule } from './app.module';

export const IS_MICROSERVICE = false;

async function bootstrap() {
  if (IS_MICROSERVICE) {
    const app = await NestFactory.createMicroservice<MicroserviceOptions>(
      AppModule,
      {
        transport: Transport.GRPC,
        options: {
          package: ['course', 'exam', 'thread'],
          protoPath: [
            __dirname + '/../../proto/course.proto',
            __dirname + '/../../proto/exam.proto',
            __dirname + '/../../proto/thread.proto',
          ],
          url: '0.0.0.0:' + process.env.COURSE_PORT,
        },
      },
    );

    app.listen();
  } else {
    const app = await NestFactory.create<NestExpressApplication>(AppModule);

    await app.listen(3000);
  }
}
bootstrap();
