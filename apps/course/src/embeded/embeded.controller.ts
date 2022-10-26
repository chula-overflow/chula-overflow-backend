import { Controller } from '@nestjs/common';
import { EmbededService } from './embeded.service';

@Controller('embeded')
export class EmbededController {
  constructor(private readonly EmbededService: EmbededService) {}
}
