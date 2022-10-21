import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MicrController } from './micr/micr.controller';
import { ThreadController } from './thread/thread.controller';

@Module({
  imports: [],
  controllers: [AppController, MicrController, ThreadController],
  providers: [AppService],
})
export class AppModule {}
