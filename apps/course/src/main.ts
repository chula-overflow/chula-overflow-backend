import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { AppModule } from './app.module';

async function bootstrap() {
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
}
bootstrap();
