import { Test, TestingModule } from '@nestjs/testing';
import { MicrController } from './micr.controller';

describe('MicrController', () => {
  let controller: MicrController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [MicrController],
    }).compile();

    controller = module.get<MicrController>(MicrController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
