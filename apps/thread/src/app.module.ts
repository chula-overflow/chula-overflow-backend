import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MicrController } from './micr/micr.controller';

@Module({
  imports: [],
  controllers: [AppController, MicrController],
  providers: [AppService],
})
export class AppModule {}
