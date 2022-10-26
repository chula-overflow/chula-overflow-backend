import { Test, TestingModule } from '@nestjs/testing';
import { EmbededController } from './embeded.controller';

describe('EmbededController', () => {
  let controller: EmbededController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [EmbededController],
    }).compile();

    controller = module.get<EmbededController>(EmbededController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
