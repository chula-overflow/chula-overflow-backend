import nlp_pb2_grpc
import nlp_pb2
import embed
import measure
import sentence_tokenizer
import grpc
from concurrent import futures
import logging

class Nlp(nlp_pb2_grpc.NlpServicer):
  def Tokenize(self, request, context):
    para = request.para
    ret = sentence_tokenizer.sentence_tokenize(para)
    return nlp_pb2.TokenizeSentences(sentences = ret)

  def Embed(self, request, context):
    sentence = request.sentence
    ret = embed.embed(sentence)
    return nlp_pb2.EncodedSentence(vector = ret)

  def Measure(self, request, context):
    v1 = request.vector1
    v2 = request.vector2
    ret = measure.cosine_sim(v1, v2)
    return nlp_pb2.MeasureSimilarity(similarity = ret)

def serve():
    port = '3003'
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    nlp_pb2_grpc.add_NlpServicer_to_server(Nlp(), server)
    server.add_insecure_port('[::]:' + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()