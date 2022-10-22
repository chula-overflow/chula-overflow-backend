import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ThreadService } from './thread.service';

@Controller('thread')
export class ThreadController {
  constructor(private readonly ThreadService: ThreadService) {}

  @GrpcMethod()
  async createProblem() {}

  @GrpcMethod()
  async upvoteThread() {}

  @GrpcMethod()
  async downvoteThread() {}

  @GrpcMethod()
  async upvoteProblem() {}

  @GrpcMethod()
  async downvoteProblem() {}
}
