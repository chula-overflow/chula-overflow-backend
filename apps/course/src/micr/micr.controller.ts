import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';

interface MicrById {
  id: number;
}

interface Micr {
  id: number;
  name: string;
}

@Controller('micr')
export class MicrController {
  @GrpcMethod()
  findOne(data: MicrById, metadata: any): Micr {
    const items = [
      { id: 1, name: 'John' },
      { id: 2, name: 'Doe' },
    ];
    return items.find(({ id }) => id === data.id);
  }
}
