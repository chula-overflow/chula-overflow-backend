import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(
    AppModule,
    {
      transport: Transport.GRPC,
      options: {
        package: 'micr',
        protoPath: __dirname + '/../../proto/thread.proto',
      },
    },
  );

  // @ts-ignore
  app.listen(() => console.log('Microservice is listening'));
}
bootstrap();
