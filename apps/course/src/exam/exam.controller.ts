import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ExamService } from './exam.service';

@Controller('exam')
export class ExamController {
  constructor(private readonly ExamService: ExamService) {}
}
