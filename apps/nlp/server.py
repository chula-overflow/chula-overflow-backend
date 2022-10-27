
import logging
import os
env = os.environ

if 'NLP_PORT' not in env:
  logging.error("'NLP_PORT' not set")
  os.abort()
port = env['NLP_PORT']

import nlp_pb2_grpc
import nlp_pb2
import embed
import measure
import sentence_tokenizer
import grpc
from concurrent import futures

class Nlp(nlp_pb2_grpc.NlpServicer):
  def Tokenize(self, request, context):
    logging.error('[Tokenize] Get req', request.para)
    para = request.para
    ret = sentence_tokenizer.sentence_tokenize(para)
    return nlp_pb2.TokenizeSentences(sentences = ret)

  def Embed(self, request, context):
    logging.error('[Embed] Get req', request.sentence)
    sentence = request.sentence
    ret = embed.embed(sentence)
    return nlp_pb2.EncodedSentence(vector = ret)

  def Measure(self, request, context):
    logging.error('[Measure] Get req', request)
    v1 = request.vector1
    v2 = request.vector2
    ret = measure.cosine_sim(v1, v2)
    return nlp_pb2.MeasureSimilarity(similarity = ret)

def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  nlp_pb2_grpc.add_NlpServicer_to_server(Nlp(), server)
  server.add_insecure_port('0.0.0.0:' + port)
  server.start()
  logging.error("Server started, listening on " + port)
  server.wait_for_termination()

if __name__ == '__main__':
  logging.basicConfig()
  serve()