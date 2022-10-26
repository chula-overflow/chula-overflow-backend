import { Test, TestingModule } from '@nestjs/testing';
import { EmbededService } from './embeded.service';

describe('EmbededService', () => {
  let service: EmbededService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [EmbededService],
    }).compile();

    service = module.get<EmbededService>(EmbededService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
