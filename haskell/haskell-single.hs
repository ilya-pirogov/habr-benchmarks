import qualified Data.Array.Base as A
import qualified Data.Array.IArray as A
import qualified Data.Array.Unboxed as A
import System.Environment

divisorSums :: Int -> Int
divisorSums n = arr A.! (n - 1)
  where arr = A.unsafeAccumArray (+) 0 (0, n - 1) premap :: A.UArray Int Int
        premap = [ (k1 * k2 - 1, if k1 /= k2 then k1 + k2 else k1)
                 | k1 <- [ 1 .. floor bound ]
                 , k2 <- [ k1 .. n `quot` k1 ]
                 ]
        bound = sqrt $ fromIntegral n :: Double

main :: IO ()
main = do
  [nStr] <- getArgs
  print $ divisorSums $ read nStr
